using Microsoft.EntityFrameworkCore;
using UserManagementAPI.Data;
using UserManagementAPI.DTOs;
using UserManagementAPI.Models;

namespace UserManagementAPI.Services;

public class UserService
{
    private readonly AppDbContext _context;

    public UserService(AppDbContext context)
    {
        _context = context;
    }

    // Get all users
    public async Task<List<User>> GetUsers()
    {
        return await _context.Users
            .OrderByDescending(x => x.CreatedAt)
            .ToListAsync();
    }

    // Get user by ID
    public async Task<User?> GetUserById(int id)
    {
        return await _context.Users
            .FirstOrDefaultAsync(x => x.Id == id);
    }

    // Create user
    public async Task<User> CreateUser(RegisterRequest request)
    {
        var user = new User
        {
            Name = request.Name,
            Email = request.Email,
            PasswordHash = BCrypt.Net.BCrypt.HashPassword(
                request.Password
            )
        };

        _context.Users.Add(user);

        await _context.SaveChangesAsync();

        return user;
    }

    // Update user
    public async Task<User?> UpdateUser(
        int id,
        UpdateUserRequest request)
    {
        var user = await GetUserById(id);

        if (user == null)
            return null;

        user.Name = request.Name;
        user.Email = request.Email;

        await _context.SaveChangesAsync();

        return user;
    }

    public async Task<User?> ChangePassword(
        int id,
        ChangePasswordRequest request)
    {
        var user = await GetUserById(id);

        if (user == null)
            return null;

         // Check current password
        bool isPasswordValid =
            BCrypt.Net.BCrypt.Verify(
                request.OldPassword,
                user.PasswordHash
            );

        if (!isPasswordValid)
            return null;


        user.PasswordHash = BCrypt.Net.BCrypt.HashPassword(request.NewPassword);

        await _context.SaveChangesAsync();

        return user;
    }

    // Delete user
    public async Task<bool> DeleteUser(int id)
    {
        var user = await GetUserById(id);

        if (user == null)
            return false;

        _context.Users.Remove(user);

        await _context.SaveChangesAsync();

        return true;
    }

    


}