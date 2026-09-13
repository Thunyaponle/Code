using Microsoft.AspNetCore.Mvc;
using UserManagementAPI.DTOs;
using UserManagementAPI.Services;

namespace UserManagementAPI.Controllers;

[ApiController]
[Route("api/users")]
public class UsersController : ControllerBase
{
    private readonly UserService _userService;

    public UsersController(UserService userService)
    {
        _userService = userService;
    }

    // GET /api/users
    [HttpGet]
    public async Task<IActionResult> GetUsers()
    {
        var users = await _userService.GetUsers();
       

        return Ok(users);
    }

    // GET /api/users/1
    [HttpGet("{id}")]
    public async Task<IActionResult> GetUser(int id)
    {
        var user = await _userService.GetUserById(id);

        if (user == null)
        {
            return NotFound(new
            {
                message = "User not found"
            });
        }

        return Ok(user);
    }

    // POST /api/users
    [HttpPost]
    public async Task<IActionResult> CreateUser(
        RegisterRequest request)
    {
        var user = await _userService.CreateUser(request);

        return CreatedAtAction(
            nameof(GetUser),
            new { id = user.Id },
            user
        );
    }

    // PUT /api/users/1
    [HttpPut("{id}")]
    public async Task<IActionResult> UpdateUser(
        int id,
        UpdateUserRequest request)
    {
        var user = await _userService
            .UpdateUser(id, request);

        if (user == null)
        {
            return NotFound(new
            {
                message = "User not found"
            });
        }

        return Ok(user);
    }

    // PUT /api/users/1/password
    [HttpPut("{id}/password")]
    public async Task<IActionResult> ChangePassword(
        int id,
        ChangePasswordRequest request)
    {
        var user = await _userService.GetUserById(id);

        if (user == null)
        {
            return NotFound(new
            {
                message = "User not found"
            });
        }

        var result = await _userService
            .ChangePassword(id, request);

        if (result == null)
        {
            return BadRequest(new
            {
                message = "Current password is incorrect"
            });
        }

        return Ok(new
        {
            message = "Password changed successfully",
            user = new
            {
                user.Id,
                user.Name,
                user.Email
            }
        });
    }

    // DELETE /api/users/1
    [HttpDelete("{id}")]
    public async Task<IActionResult> DeleteUser(int id)
    {
        var result = await _userService.DeleteUser(id);

        if (!result)
        {
            return NotFound(new
            {
                message = "User not found"
            });
        }

        return NoContent();
    }
}